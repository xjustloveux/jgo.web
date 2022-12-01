export class JobLogClass {
  name: string;
  id: string;
  ct_date: Date;

  constructor(data: any) {
    this.name = data.name;
    this.id = data.id;
    this.ct_date = data.ct_date;
  }
}
