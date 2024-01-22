import {Component, Input} from '@angular/core';

import {LogFileClass} from '../../model/logFileClass';

@Component({
  selector: 'app-log-file',
  templateUrl: './log-file.component.html',
})
export class LogFileComponent {

  @Input('content') content: LogFileClass | undefined = undefined;
}
