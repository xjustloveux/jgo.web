export class LogFileClass {
  name: string;
  dir: boolean
  child: LogFileClass[];

  constructor(data: any) {
    this.name = data.name;
    this.dir = data.dir;
    this.child = data.child;
  }
}
