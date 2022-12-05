import {MsgClass} from "./msgClass";

export class ResMsgClass {
  rows: MsgClass[];
  rowStart: number;
  rowEnd: number;
  totalRecord: number;

  constructor(data: any) {
    this.rows = data.rows;
    this.rowStart = data.rowStart;
    this.rowEnd = data.rowEnd;
    this.totalRecord = data.totalRecord;
  }
}
