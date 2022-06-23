import {Component, OnInit} from '@angular/core';
import {B103Service} from './b103.service';
import {PkgContentClass} from '../../../model/pkgContentClass';
import Swal from 'sweetalert2';
import {LogFileClass} from "../../../model/logFileClass";

@Component({
  selector: 'app-b103',
  templateUrl: './b103.component.html',
  styleUrls: ['./b103.component.css']
})
export class B103Component implements OnInit {

  proId = 'B103';
  content: PkgContentClass[] = [];
  log: LogFileClass | undefined;

  constructor(
    public service: B103Service,
  ) {
    this.service.queryContent(this.proId).then((res) => {
      if (res.success) {
        this.content = res.content;
      } else {
        Swal.fire('Query content failed!');
      }
    });
    this.service.queryLog().then((res) => {
      if (res.success) {
        this.log = res.list;
      } else {
        Swal.fire('Query log failed!');
      }
    });
  }

  ngOnInit(): void {
  }
}
