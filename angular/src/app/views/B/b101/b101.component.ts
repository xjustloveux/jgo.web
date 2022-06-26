import {Component, OnInit} from '@angular/core';
import {Title} from '@angular/platform-browser';
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
    public title: Title
  ) {
    this.title.setTitle('jsql | JGo');
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
