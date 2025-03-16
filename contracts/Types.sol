// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {SD59x18} from "@prb/math/src/SD59x18.sol";

library Types {
    struct NetworkState {
        uint deployedAt; // Unix time
    }

    struct NetworkConfig {
        uint64 epochTime; // In seconds
        uint64 searchRangeTime; // In seconds
        SD59x18 k_evn; // k_EVN
        SD59x18 pf; // pf
        SD59x18 y; // γ
        SD59x18 k_rup; // k_RUP
        SD59x18 w_c; // w_C
    }

    enum MetricType {
        Undefined,
        Ascendant, // Best values tend to +inf
        Descendant, // Best values tent to 0
        Alert // Alert-based metrics
    }

    struct TrustNodeData {
        address mAddr; // Trust manager address
        address nAddr; // Trust node address (owner)
        string ip; // Trust node IP
        uint64 nextSLAId; // Starting at 0
        Types.Evidence[] evidences; // Evidences received by other trust nodes
        uint registeredAt; // Unix time
    }

    struct SLA {
        address customer;
        SLAMetric[] metrics;
        uint setAt; // Unix time
        uint acceptedAt; // Unix time
        uint terminatedAt; // Unix time
    }

    struct Evidence {
        EvidenceInfo evi;
        mapping(string => bool) metricIndex; // Metric ID => repeated metric?
        uint sentAt; // Unix time
    }

    struct EvidenceInfo {
        uint64 slaId;
        address sender;
        EvidenceMetric[] metrics;
    }

    struct SLAMetric {
        string id; // Metric ID
        SD59x18 threshold; // T_m
        SD59x18 weighting; // w_m
    }

    struct EvidenceMetric {
        string id; // Metric ID
        SD59x18 value; // x_m
    }

    enum DTSType {
        a, // α
        b // β
    }

    struct TrustData {
        mapping(address => DTSInfo) dtsis; // Target EOA => direct trust score info
        RInfo[] ris; // Array of historical Rs
    }

    // Direct score info
    struct DTSInfo {
        SD59x18 a; // α
        SD59x18 b; // β
        SD59x18 dts; // DTS
        SD59x18 r; // TODO (R_j)
        SD59x18 cp; // TODO (C')
        uint64 na; // Nα (total count)
        uint64 nb; // Nβ (total count)
        NCount[] nCounts;
    }

    struct NCount {
        uint64 na; // Nα (count per evidence)
        uint64 nb; // Nβ (count per evidence)
        uint countedAt; // Unix time
    }

    struct RInfo {
        SD59x18 r; // R_i
        uint computedAt; // Unix time
    }
}
