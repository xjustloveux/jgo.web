import {Component, OnInit} from '@angular/core';
import {B101Service} from "./b101.service";
import {PkgContentClass} from '../../../model/pkgContentClass';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-b101',
  templateUrl: './b101.component.html',
  styleUrls: ['./b101.component.css']
})
export class B101Component implements OnInit {

  proId = 'B101';
  content: PkgContentClass[] = [];

  constructor(
    public service: B101Service,
  ) {
    this.service.queryContent(this.proId).then((res) => {
      if (res.success) {
        this.content = res.content;
      } else {
        Swal.fire('Query content failed!');
      }
    });
  }

  ngOnInit(): void {
  }
}
