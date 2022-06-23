import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {B101Component} from './b101.component';
import {PkgContentModule} from '../../../shared/pkg-content/pkg-content.module';

@NgModule({
  declarations: [B101Component],
  imports: [CommonModule, PkgContentModule]
})
export class B101Module {
}
