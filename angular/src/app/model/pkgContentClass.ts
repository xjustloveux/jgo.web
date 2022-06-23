export class PkgContentClass {
  Title: string;
  Description: string;
  Code: string;
  Note: string;

  constructor(data: any) {
    this.Title = data.Title;
    this.Description = data.Description;
    this.Code = data.Code;
    this.Note = data.Note;
  }
}
