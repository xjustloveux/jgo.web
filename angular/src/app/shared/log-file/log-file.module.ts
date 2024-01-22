import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {LogFileComponent} from './log-file.component';

@NgModule({
  declarations: [LogFileComponent],
  exports: [LogFileComponent],
  imports: [CommonModule]
})
export class LogFileModule {
}
