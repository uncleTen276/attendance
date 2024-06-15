import { ethers } from "hardhat";
import fs from "fs";
import path from "path";
async function main() {
  const Attendance = await ethers.getContractFactory("Attendance");
  const contract = await Attendance.deploy();
  const address = await contract.getAddress();
  // write to env file
  const envFilePath = path.resolve(__dirname, "..", "..", "server", ".env");
  let envConfig = fs.existsSync(envFilePath)
    ? fs.readFileSync(envFilePath, "utf8")
    : "";
  const newContent = `CONTRACT_ADDRESS=${address}\n`;
  if (envConfig.includes(`CONTRACT_ADDRESS=`)) {
    envConfig = envConfig.replace(
      `CONTRACT_ADDRESS=`,
      `CONTRACT_ADDRESS=${newContent}`
    );
  } else {
    envConfig += newContent;
  }

  fs.writeFileSync(envFilePath, envConfig, "utf8");
  console.log(`Contract deployed to ${address}`);
}

main().catch((error) => {
  console.log(error);
  process.exitCode = 1;
});
