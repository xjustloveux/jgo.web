export class JobLogClass {
  Name: string;
  Id: string;
  CtDate: Date;

  constructor(data: any) {
    this.Name = data.Name;
    this.Id = data.Id;
    this.CtDate = data.CtDate;
  }
}
