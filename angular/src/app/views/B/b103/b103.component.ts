import {Component, OnInit} from '@angular/core';
import {Title} from '@angular/platform-browser';
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

  proId = 'b103';
  content: PkgContentClass[] = [];
  log: LogFileClass | undefined;

  constructor(
    public service: B103Service,
    public title: Title
  ) {
    this.title.setTitle('jlog | JGo');
    this.service.queryContent(this.proId).then((res) => {
      this.content = res;
    });
    this.service.queryLog().then((res) => {
      this.log = res;
    });
  }

  ngOnInit(): void {
  }
}
