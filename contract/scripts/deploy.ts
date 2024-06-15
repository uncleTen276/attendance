import { ethers } from "hardhat";

const ether = require("hardhat").ethers;

async function main() {
  const Attendance = await ethers.getContractFactory("Attendance");
  const contract = await Attendance.deploy();
  const address = await contract.getAddress();
  console.log(`Contract Address, ${address}`);
}

main().catch((error) => {
  console.log(error);
  process.exitCode = 1;
});
