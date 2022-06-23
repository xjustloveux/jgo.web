import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {B101Module} from './b101/b101.module';
import {B102Module} from './b102/b102.module';
import {B103Module} from './b103/b103.module';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    B101Module,
    B102Module,
    B103Module,
  ]
})
export class BModule {
}
