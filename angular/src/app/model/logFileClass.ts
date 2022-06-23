export class LogFileClass {
  Name: string;
  Dir: boolean
  Child: LogFileClass[];

  constructor(data: any) {
    this.Name = data.Name;
    this.Dir = data.Dir;
    this.Child = data.Child;
  }
}
