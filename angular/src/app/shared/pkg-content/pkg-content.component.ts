import {Component, OnInit, Input} from '@angular/core';
import {PkgContentClass} from '../../model/pkgContentClass';

@Component({
  selector: 'app-pkg-content',
  templateUrl: './pkg-content.component.html',
  styleUrls: ['./pkg-content.component.css']
})
export class PkgContentComponent implements OnInit {

  @Input('content') content: PkgContentClass[] = [];

  constructor() {
  }

  ngOnInit(): void {
  }
}
