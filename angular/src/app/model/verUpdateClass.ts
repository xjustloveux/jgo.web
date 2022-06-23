export class VerUpdateClass {
  Seq: number;
  VerS: number;
  VerE: number;
  Content: string;

  constructor(data: any) {
    this.Seq = data.Seq;
    this.VerS = data.VerS;
    this.VerE = data.VerE;
    this.Content = data.Content;
  }
}
