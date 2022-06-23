import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {B102Component} from './b102.component';
import {PkgContentModule} from '../../../shared/pkg-content/pkg-content.module';

@NgModule({
  declarations: [B102Component],
  imports: [
    CommonModule,
    PkgContentModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class B102Module {
}
