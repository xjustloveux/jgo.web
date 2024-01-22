import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";

import {A101Component} from './a101.component';

@NgModule({
  declarations: [A101Component],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class A101Module {
}
