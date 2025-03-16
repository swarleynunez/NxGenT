// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

//import "hardhat/console.sol";
import "./Types.sol";
import "./TrustNode.sol";

contract TrustManager {
    // Network state and configuration
    Types.NetworkState public state;
    Types.NetworkConfig public config;

    // SLA metrics
    mapping(string => Types.MetricType) private metrics; // Metric ID => metric type

    // Trust nodes
    mapping(address => address) public nodes; // Trust node EOA => trust node contract
    mapping(string => address) public nodeIPs; // Trust node IP => trust node EOA
    address[] private registeredNodes; // Trust node EOAs

    // Trust data (direct, indirect and reputation scores)
    mapping(address => Types.TrustData) private trustData; // Trust node EOA => trust data

    constructor(uint64 epTime, uint64 srTime) {
        // Initialize network
        state = Types.NetworkState(block.timestamp);
        config = Types.NetworkConfig(
            epTime,
            srTime,
            sd(10e18), // k_EVN
            sd(3e18), // pf
            sd(10e18), // γ
            sd(100e18), // k_RUP
            sd(0.25e18) // w_C
        );

        // Initialize SLA metrics
        metrics["availability"] = Types.MetricType.Ascendant;
        metrics["latency"] = Types.MetricType.Descendant;
        metrics["jitter"] = Types.MetricType.Descendant;
        metrics["cpu_cores"] = Types.MetricType.Ascendant;
        metrics["cpu_frequency"] = Types.MetricType.Ascendant;
        metrics["memory_total"] = Types.MetricType.Ascendant;
        metrics["disk_total"] = Types.MetricType.Ascendant;
        metrics["ids_alert"] = Types.MetricType.Alert;
    }

    ///////////////
    /// Setters ///
    ///////////////
    function registerTrustNode(string memory ip) external {
        require(
            !isTrustNodeRegistered(msg.sender),
            "Trust node already registered"
        );
        require(isIPAvailable(ip), "IP not available");

        // Create a new trust node
        nodes[msg.sender] = address(new TrustNode(msg.sender, ip));
        nodeIPs[ip] = msg.sender;
        registeredNodes.push(msg.sender);
        trustData[msg.sender].ris.push(
            Types.RInfo(sd(0.5e18), block.timestamp)
        );
    }

    /*function updateTrustNodeIP(string memory ip) external {
        require(isTrustNodeRegistered(msg.sender), "Trust node not registered");
        require(isIPAvailable(ip), "IP not available");

        // Update state
        nodeIPs[TrustNode(nodes[msg.sender]).getIP()] = address(0);
        nodeIPs[ip] = msg.sender;
        TrustNode(nodes[msg.sender]).updateIP(ip);
    }*/

    function sendEvidence(
        address target,
        uint64 _slaId,
        Types.EvidenceMetric[] memory _metrics
    ) external {
        require(msg.sender != target, "Self-evidences not allowed");
        require(isTrustNodeRegistered(msg.sender), "Sender not registered");
        require(isTrustNodeRegistered(target), "Target not registered");
        require(
            TrustNode(nodes[target]).isSLACustomer(_slaId, msg.sender),
            "Sender is not the SLA customer"
        );
        require(TrustNode(nodes[target]).isSLAActive(_slaId), "SLA not active");
        /*require(
            canSendEvidence(msg.sender, target, _slaId),
            "Sender cannot submit an evidence yet"
        );*/
        require(_metrics.length > 0, "No metrics specified");

        // Create a new evidence
        TrustNode(nodes[target]).storeEvidence(_slaId, _metrics);
        TrustNode(nodes[msg.sender]).updateLastEvidenceTime(target, _slaId);

        // Get direct trust score info
        Types.DTSInfo storage dtsi = trustData[msg.sender].dtsis[target];

        // Compute evidence metrics
        uint64 evNa;
        uint64 evNb;
        for (uint i = 0; i < _metrics.length; i++) {
            // Get SLA metric info
            Types.MetricType mt = metrics[_metrics[i].id];
            (SD59x18 t_m, SD59x18 w_m) = TrustNode(nodes[target]).getSLAMetric(
                _slaId,
                _metrics[i].id
            );

            // Check x_m and get xp_m
            (SD59x18 xp_m, Types.DTSType dtst) = checkEvidenceMetric(
                mt,
                _metrics[i].value,
                t_m
            );

            // Update α/β and count metric
            if (dtst == Types.DTSType.a) {
                dtsi.a = dtsi.a.add(xp_m.mul(w_m));
                dtsi.na++;
                evNa++;
            } else if (dtst == Types.DTSType.b) {
                dtsi.b = dtsi.b.add(xp_m.mul(w_m));
                dtsi.nb++;
                evNb++;
            }
        }

        // Update direct trust score info
        dtsi.dts = computeNewDTS(dtsi.a, dtsi.b); // Update DTS
        dtsi.nCounts.push(Types.NCount(evNa, evNb, block.timestamp)); // Create a new N count

        // Create a new R
        trustData[target].ris.push(
            Types.RInfo(computeNewRScore(target), block.timestamp)
        );
    }

    function checkEvidenceMetric(
        Types.MetricType mt,
        SD59x18 x_m,
        SD59x18 t_m
    ) public view returns (SD59x18 xp_m, Types.DTSType dtst) {
        // Check metric type
        if (mt == Types.MetricType.Ascendant) {
            if (x_m.gte(t_m)) {
                // Contributes positively
                xp_m = sd(1e18);
                dtst = Types.DTSType.a;

                // Experiments
                //xp_m = normalizeMetricValue(x_m, t_m);
            } else {
                // Contributes negatively
                xp_m = normalizeMetricValue(x_m, t_m);
                dtst = Types.DTSType.b;
            }
        } else if (mt == Types.MetricType.Descendant) {
            if (x_m.lte(t_m)) {
                // Contributes positively
                xp_m = sd(1e18);
                dtst = Types.DTSType.a;

                // Experiments
                //xp_m = normalizeMetricValue(x_m, t_m);
            } else {
                // Contributes negatively
                xp_m = normalizeMetricValue(x_m, t_m);
                dtst = Types.DTSType.b;
            }
        } else if (mt == Types.MetricType.Alert) {
            // Contributes negatively
            xp_m = sd(1e18);
            dtst = Types.DTSType.b;
        }
    }

    function normalizeMetricValue(
        SD59x18 x_m,
        SD59x18 t_m
    ) public view returns (SD59x18) {
        return
            t_m
                .sub(x_m)
                .abs()
                .div(config.k_evn)
                .mul(sd(-1e18))
                .exp()
                .sub(sd(1e18))
                .abs();
    }

    function computeNewDTS(SD59x18 a, SD59x18 b) public view returns (SD59x18) {
        return a.div(b.mul(config.pf).add(a).add(config.y));
    }

    function computeNewRScore(address target) public returns (SD59x18 r) {
        SD59x18 nc; // Numerator counter
        SD59x18 dc; // Denominator counter
        // Experiments
        //uint64 dc; // Denominator counter
        for (uint i = 0; i < registeredNodes.length; i++) {
            // Check if node has direct trust score for the target node
            if (!hasDirectScore(registeredNodes[i], target)) {
                continue;
            }

            // Get direct trust score info
            Types.DTSInfo storage dtsi = trustData[registeredNodes[i]].dtsis[
                target
            ];

            // Get necessary fields
            SD59x18 r_j = getRScore(registeredNodes[i]); // R_j
            SD59x18 cp_ji = computeC(
                registeredNodes[i],
                target,
                dtsi.na,
                dtsi.nb
            ); // C'_(j,i)

            dtsi.r = r_j; // TODO
            dtsi.cp = cp_ji; // TODO

            // Update counters
            nc = nc.add(dtsi.dts.mul(r_j).mul(cp_ji));
            dc = dc.add(r_j.mul(cp_ji));

            // Experiments
            //nc = nc.add(dtsi.dts);
            //dc++;
        }

        if (!dc.isZero()) r = nc.div(dc);
        else r = sd(0);

        // Experiments
        //if (dc > 0) r = nc.div(convert(int256(int64(dc))));
        //else r = sd(0);
    }

    function computeC(
        address addr,
        address target,
        uint64 na,
        uint64 nb
    ) public view returns (SD59x18) {
        // Get N count info
        (uint64 na_in, uint64 nb_in) = getCurrentIntervalNCount(addr, target);
        SD59x18 n_in = convert(int256(int64(na_in + nb_in)));
        SD59x18 n_out = convert(int256(int64(na - na_in + nb - nb_in)));

        // Compute C applying forgetting factor w_C
        SD59x18 c = sd(1e18).sub(config.w_c).mul(n_in).add(
            n_out.mul(config.w_c)
        );

        return c.div(config.k_rup).mul(sd(-1e18)).exp().sub(sd(1e18)).abs();
    }

    ///////////////
    /// Getters ///
    ///////////////
    function getRegisteredTrustNodes()
        external
        view
        returns (address[] memory)
    {
        return registeredNodes;
    }

    function getDTSInfo(
        address addr,
        address target
    ) external view returns (Types.DTSInfo memory) {
        return trustData[addr].dtsis[target];
    }

    function getRScore(address addr) public view returns (SD59x18 r) {
        Types.RInfo[] storage ris = trustData[addr].ris;
        /*for (uint i = ris.length; i >= 1; i--) {
            r = ris[i - 1].r;
        }*/
        r = ris[ris.length - 1].r;
    }

    function getRScoreCount(address addr) external view returns (uint) {
        return trustData[addr].ris.length;
    }

    function getCurrentIntervalNCount(
        address addr,
        address target
    ) public view returns (uint64 na_in, uint64 nb_in) {
        Types.NCount[] storage nCounts = trustData[addr].dtsis[target].nCounts;
        for (uint i = nCounts.length; i >= 1; i--) {
            if (
                block.timestamp - config.searchRangeTime <=
                nCounts[i - 1].countedAt &&
                nCounts[i - 1].countedAt <= block.timestamp
            ) {
                na_in = na_in + nCounts[i - 1].na;
                nb_in = nb_in + nCounts[i - 1].nb;
            } else {
                break;
            }
        }
    }

    ///////////////
    /// Helpers ///
    ///////////////
    function isTrustNodeRegistered(address addr) public view returns (bool) {
        if (nodes[addr] != address(0)) return true;
        return false;
    }

    function isIPAvailable(string memory ip) public view returns (bool) {
        if (nodeIPs[ip] == address(0)) return true;
        return false;
    }

    function canSendEvidence(
        address sender,
        address target,
        uint64 slaId
    ) public view returns (bool) {
        if (
            TrustNode(nodes[sender]).getLastEvidenceTime(target, slaId) +
                config.epochTime <=
            block.timestamp
        ) return true;
        return false;
    }

    function existMetric(string memory metricId) public view returns (bool) {
        if (metrics[metricId] != Types.MetricType.Undefined) return true;
        return false;
    }

    function hasDirectScore(
        address addr,
        address target
    ) public view returns (bool) {
        if (
            trustData[addr].dtsis[target].na > 0 ||
            trustData[addr].dtsis[target].nb > 0
        ) return true;
        return false;
    }
}
