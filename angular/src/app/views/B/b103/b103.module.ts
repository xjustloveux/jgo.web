import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {B103Component} from './b103.component';
import {PkgContentModule} from '../../../shared/pkg-content/pkg-content.module';
import {LogFileModule} from '../../../shared/log-file/log-file.module';

@NgModule({
  declarations: [B103Component],
  imports: [CommonModule, PkgContentModule, LogFileModule]
})
export class B103Module {
}
