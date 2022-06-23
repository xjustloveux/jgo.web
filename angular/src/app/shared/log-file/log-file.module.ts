import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {NgxBootstrapIconsModule, allIcons} from 'ngx-bootstrap-icons';
import {LogFileComponent} from './log-file.component';

@NgModule({
  declarations: [LogFileComponent],
  exports: [LogFileComponent],
  imports: [
    CommonModule,
    NgxBootstrapIconsModule.pick(allIcons)
  ]
})
export class LogFileModule {
}
