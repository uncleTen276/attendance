// SPDX-License-Identifier: MIT
pragma experimental ABIEncoderV2;
pragma solidity >=0.4.0 <=0.9.0;
contract Owned{
  address owner;
  constructor() {
    owner = msg.sender;
  }

  modifier onlyOwner{
    require(msg.sender == owner, "Only owner can call this function");
    _;
  }

}
contract Attendance is Owned{
  struct EmployeeRecord{
      uint256 recordId;
      uint256 employeeId ;
      string checkInTime;
      string position ;
      string details;
  }

  uint256 public nextRecordId;

  mapping (uint256 => EmployeeRecord) public attendanceRecords;
  mapping(uint256 => uint256[]) public employeeRecordIds;

  function recordAtttendance(uint256 employeeId, string memory checkInTime, string memory details, string memory position) public onlyOwner returns (uint256){
    uint256 recordId = nextRecordId;
    attendanceRecords[nextRecordId] = EmployeeRecord({
      recordId: recordId,
      employeeId: employeeId,
      checkInTime: checkInTime,
      details: details,
      position: position
    });
    employeeRecordIds[employeeId].push(nextRecordId);
    nextRecordId++;
    return recordId;
  }

   function getAttendanceByEmployee(uint256 _employeeId) public view returns (EmployeeRecord[] memory) {
        uint256[] memory recordIds = employeeRecordIds[_employeeId];
        EmployeeRecord[] memory employeeRecords = new EmployeeRecord[](recordIds.length);

        for (uint256 i = 0; i < recordIds.length; i++) {
            employeeRecords[i] = attendanceRecords[recordIds[i]];
        }

        return employeeRecords;
    }

     function updateAttendance(
        uint256 _recordId,
        string memory _checkInTime,
        string memory _details,
        string memory _position
    ) public onlyOwner {
        require(attendanceRecords[_recordId].recordId == _recordId, "Record does not exist");

        EmployeeRecord storage record = attendanceRecords[_recordId];

        record.checkInTime = _checkInTime;
        record.details = _details;
        record.position = _position;
    }}
