export class PkgContentClass {
  title: string;
  description: string;
  code: string;
  note: string;

  constructor(data: any) {
    this.title = data.title;
    this.description = data.description;
    this.code = data.code;
    this.note = data.note;
  }
}
