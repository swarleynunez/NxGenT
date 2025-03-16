// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "./Types.sol";
import "./TrustManager.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import {sd, convert} from "@prb/math/src/SD59x18.sol";

contract TrustNode {
    // Object containing trust node data
    Types.TrustNodeData public data;

    // Service Level Agreements offered to other trust nodes
    mapping(uint64 => Types.SLA) public SLAs; // SLA ID => SLA

    // Last evidence time per target and SLA
    mapping(address => mapping(uint64 => uint)) public lastEvidenceTimes; // Target EOA => (SLA ID => unix time)

    // Events
    event NewSLA(address from, uint64 id, address to); // Provider EOA, SLA ID, customer EOA
    event AcceptedSLA(address from, uint64 id, address to); // Provider EOA, SLA ID, customer EOA
    event TerminatedSLA(address from, uint64 id, address to); // Provider EOA, SLA ID, customer EOA

    constructor(address _nAddr, string memory _ip) {
        data.mAddr = msg.sender;
        data.nAddr = _nAddr;
        data.ip = _ip;
        data.registeredAt = block.timestamp;
    }

    modifier onlyTrustManager() {
        require(
            msg.sender == data.mAddr,
            "Only trust manager can call this function"
        );
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == data.nAddr, "Only owner can call this function");
        _;
    }

    ///////////////
    /// Setters ///
    ///////////////
    /*function updateIP(string memory _ip) external onlyTrustManager {
        data.ip = _ip;
    }*/

    function setSLA(
        address _customer,
        Types.SLAMetric[] memory _metrics
    ) external onlyOwner {
        require(msg.sender != _customer, "Self-SLAs not allowed");
        require(
            TrustManager(data.mAddr).isTrustNodeRegistered(_customer),
            "Customer not registered"
        );
        require(_metrics.length > 0, "No metrics specified");

        // Create a new SLA
        Types.SLA storage sla = SLAs[data.nextSLAId];
        sla.customer = _customer;
        sla.setAt = block.timestamp;
        for (uint i = 0; i < _metrics.length; i++) {
            require(
                TrustManager(data.mAddr).existMetric(_metrics[i].id),
                "Metric not found"
            );
            require(
                !hasSLAMetric(data.nextSLAId, _metrics[i].id),
                "Repeated metric"
            );

            // Store SLA metric
            sla.metrics.push(_metrics[i]);
        }

        // Send event
        emit NewSLA(msg.sender, data.nextSLAId, _customer);

        // Update state
        data.nextSLAId++;
    }

    function acceptSLA(uint64 id) external {
        require(canAcceptSLA(id, msg.sender), "Cannot accept SLA");

        // Update state
        SLAs[id].acceptedAt = block.timestamp;

        // Send event
        emit AcceptedSLA(data.nAddr, id, SLAs[id].customer);
    }

    /*function terminateSLA(uint64 id) external {
        require(canTerminateSLA(id, msg.sender), "Cannot terminate SLA");

        // Update state
        SLAs[id].terminatedAt = block.timestamp;

        // Send event
        emit TerminatedSLA(data.nAddr, id, SLAs[id].customer);
    }*/

    function storeEvidence(
        uint64 _slaId,
        Types.EvidenceMetric[] memory _metrics
    ) external onlyTrustManager {
        Types.Evidence storage ev = data.evidences.push();
        ev.evi.slaId = _slaId;
        ev.evi.sender = tx.origin;
        ev.sentAt = block.timestamp;
        for (uint i = 0; i < _metrics.length; i++) {
            require(
                TrustManager(data.mAddr).existMetric(_metrics[i].id),
                "Metric not found"
            );
            require(
                hasSLAMetric(_slaId, _metrics[i].id),
                "Metric not found in SLA"
            );
            require(!ev.metricIndex[_metrics[i].id], "Repeated metric");
            ev.metricIndex[_metrics[i].id] = true;

            ev.evi.metrics.push(_metrics[i]);
        }
    }

    function updateLastEvidenceTime(
        address target,
        uint64 slaId
    ) external onlyTrustManager {
        lastEvidenceTimes[target][slaId] = block.timestamp;
    }

    ///////////////
    /// Getters ///
    ///////////////
    function getIP() external view returns (string memory) {
        return data.ip;
    }

    function getLastEvidenceTime(
        address target,
        uint64 slaId
    ) external view returns (uint) {
        return lastEvidenceTimes[target][slaId];
    }

    function getEvidencesCount() external view returns (uint) {
        return data.evidences.length;
    }

    function getSLAMetric(
        uint64 slaId,
        string memory metricId
    ) external view returns (SD59x18, SD59x18) {
        Types.SLAMetric[] storage metrics = SLAs[slaId].metrics;
        for (uint i = 0; i < metrics.length; i++) {
            if (Strings.equal(metrics[i].id, metricId))
                return (metrics[i].threshold, metrics[i].weighting);
        }

        return (sd(0), sd(0));
    }

    /*function getSLAMetrics(
        uint64 slaId
    ) external view returns (Types.SLAMetric[] memory) {
        return (SLAs[slaId].metrics);
    }*/

    ///////////////
    /// Helpers ///
    ///////////////
    function isSLACustomer(uint64 id, address addr) public view returns (bool) {
        if (addr == SLAs[id].customer) return true;
        return false;
    }

    function isSLAActive(uint64 id) public view returns (bool) {
        if (SLAs[id].acceptedAt > 0 && SLAs[id].terminatedAt == 0) return true;
        return false;
    }

    function canAcceptSLA(uint64 id, address addr) public view returns (bool) {
        if (
            addr == SLAs[id].customer &&
            SLAs[id].acceptedAt == 0 &&
            SLAs[id].terminatedAt == 0
        ) return true;
        return false;
    }

    /*function canTerminateSLA(
        uint64 id,
        address addr
    ) public view returns (bool) {
        if (
            (addr == data.nAddr || addr == SLAs[id].customer) &&
            SLAs[id].terminatedAt == 0
        ) return true;
        return false;
    }*/

    function hasSLAMetric(
        uint64 slaId,
        string memory metricId
    ) public view returns (bool) {
        for (uint i = 0; i < SLAs[slaId].metrics.length; i++) {
            if (Strings.equal(SLAs[slaId].metrics[i].id, metricId)) return true;
        }

        return false;
    }
}
