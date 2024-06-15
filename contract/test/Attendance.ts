import { expect } from "chai";
import { ethers } from "hardhat";
import { Attendance, Attendance__factory } from "../typechain-types";

describe("Attendance Contract", function () {
  let attendance: Attendance;
  let owner: any;
  let addr1: any;

  beforeEach(async function () {
    // Get the ContractFactory and Signers here.
    const AttendanceFactory: Attendance__factory =
      await ethers.getContractFactory("Attendance");
    [owner, addr1] = await ethers.getSigners();

    // Deploy the contract with the owner as the deployer
    attendance = await AttendanceFactory.deploy();
  });

  it("Should record attendance by the owner", async function () {
    const employeeId = 1;
    const checkInTime = "2023-06-14 09:00:00";
    const details = "Checked in on time.";
    const position = "Developer";

    const tx = await attendance.recordAtttendance(
      employeeId,
      checkInTime,
      details,
      position
    );
    await tx.wait();

    const record = await attendance.attendanceRecords(0);

    expect(record.recordId).to.equal(0);
    expect(record.employeeId).to.equal(employeeId);
    expect(record.checkInTime).to.equal(checkInTime);
    expect(record.details).to.equal(details);
    expect(record.position).to.equal(position);
  });

  it("Should not allow non-owner to record attendance", async function () {
    const employeeId = 1;
    const checkInTime = "2023-06-14 09:00:00";
    const details = "Checked in on time.";
    const position = "Developer";

    await expect(
      attendance
        .connect(addr1)
        .recordAtttendance(employeeId, checkInTime, details, position)
    ).to.be.revertedWith("Only owner can call this function");
  });

  it("Should return attendance records by employee", async function () {
    const employeeId = 1;
    const checkInTime1 = "2023-06-14 09:00:00";
    const details1 = "Checked in on time.";
    const position1 = "Developer";

    const checkInTime2 = "2023-06-15 09:05:00";
    const details2 = "Checked in 5 minutes late.";
    const position2 = "Developer";

    await attendance.recordAtttendance(
      employeeId,
      checkInTime1,
      details1,
      position1
    );
    await attendance.recordAtttendance(
      employeeId,
      checkInTime2,
      details2,
      position2
    );

    const records = await attendance.getAttendanceByEmployee(employeeId);

    expect(records.length).to.equal(2);
    expect(records[0].checkInTime).to.equal(checkInTime1);
    expect(records[1].checkInTime).to.equal(checkInTime2);
  });

  it("Should allow owner to update attendance record", async function () {
    const employeeId = 1;
    const checkInTime = "2023-06-14 09:00:00";
    const details = "Checked in on time.";
    const position = "Developer";

    const tx = await attendance.recordAtttendance(
      employeeId,
      checkInTime,
      details,
      position
    );
    await tx.wait();

    const updatedCheckInTime = "2023-06-14 09:30:00";
    const updatedDetails = "Checked in late.";
    const updatedPosition = "Senior Developer";

    const updateTx = await attendance.updateAttendance(
      0,
      updatedCheckInTime,
      updatedDetails,
      updatedPosition
    );
    await updateTx.wait();

    const updatedRecord = await attendance.attendanceRecords(0);

    expect(updatedRecord.checkInTime).to.equal(updatedCheckInTime);
    expect(updatedRecord.details).to.equal(updatedDetails);
    expect(updatedRecord.position).to.equal(updatedPosition);
  });

  it("Should not allow non-owner to update attendance record", async function () {
    const employeeId = 1;
    const checkInTime = "2023-06-14 09:00:00";
    const details = "Checked in on time.";
    const position = "Developer";

    const tx = await attendance.recordAtttendance(
      employeeId,
      checkInTime,
      details,
      position
    );
    await tx.wait();

    const updatedCheckInTime = "2023-06-14 09:30:00";
    const updatedDetails = "Checked in late.";
    const updatedPosition = "Senior Developer";

    await expect(
      attendance
        .connect(addr1)
        .updateAttendance(
          0,
          updatedCheckInTime,
          updatedDetails,
          updatedPosition
        )
    ).to.be.revertedWith("Only owner can call this function");
  });
});
