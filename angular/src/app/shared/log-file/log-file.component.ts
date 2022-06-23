import {Component, OnInit, Input} from '@angular/core';
import {LogFileClass} from '../../model/logFileClass';

@Component({
  selector: 'app-log-file',
  templateUrl: './log-file.component.html',
  styleUrls: ['./log-file.component.css']
})
export class LogFileComponent implements OnInit {

  @Input('content') content: LogFileClass | undefined = undefined;

  constructor() {
  }

  ngOnInit(): void {
  }
}
