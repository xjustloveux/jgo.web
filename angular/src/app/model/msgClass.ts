export class MsgClass {
  content: string;
  ct_date: Date;

  constructor(data: any) {
    this.content = data.content;
    this.ct_date = data.ct_date;
  }
}
