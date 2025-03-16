import {ethers} from "hardhat";
import {TrustManager} from "../typechain-types";

describe("TrustManager", function () {
    let tmContract: TrustManager;
    let accounts: any;

    before(async function () {
        tmContract = await ethers.deployContract("TrustManager", [5]);
        accounts = await ethers.getSigners();
    });

    /*it("Test 1. Registrations", async function () {
        await tmContract.registerTrustNode("192.168.0.1");
        await tmContract.connect(accounts[1]).registerTrustNode("192.168.0.2");
        await tmContract.connect(accounts[2]).registerTrustNode("192.168.0.3");
    });

    it("Test 2. SLAs", async function () {
        const tnContract = await ethers.getContractAt("TrustNode", await tmContract.nodes(accounts[2]));
        await tnContract.connect(accounts[2]).setSLA(accounts[0], "{DATA}");
        await tnContract.connect(accounts[0]).acceptSLA();
    });

    it("Test 3. Evidences", async function () {
        let n = 5;
        while (n > 0) {
            await tmContract.sendEvidence(accounts[2], "{DATA}");
            await time.increase(4);
            await new Promise(f => setTimeout(f, 1000));
            n--;
        }
    });

    it("Test 4. Data", async function () {
        const evidences = await tmContract.getEvidences(accounts[2]);
        console.log(evidences);

        const tnContract = await ethers.getContractAt("TrustNode", await tmContract.nodes(accounts[2]));
        console.log(await tnContract.getIP(), await tnContract.getLastEvidenceTime());
        console.log(await tnContract.hasSLA(accounts[0]), await tnContract.isSLAAccepted(accounts[0]));
    });

    it("Test 5. SLA termination", async function () {
        const tnContract = await ethers.getContractAt("TrustNode", await tmContract.nodes(accounts[2]));
        await tnContract.connect(accounts[0]).terminateSLA();
    });*/
});
