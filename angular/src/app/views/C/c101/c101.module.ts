import {NgModule, Pipe, PipeTransform} from '@angular/core';
import {CommonModule} from '@angular/common';
import {DomSanitizer, SafeHtml} from '@angular/platform-browser';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {C101Component} from './c101.component';

@Pipe({name: 'sanitizeHtml'})
export class SanitizeHtmlPipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) {
  }

  transform(v: string): SafeHtml {
    return this.sanitizer.bypassSecurityTrustHtml(v);
  }
}

@NgModule({
  declarations: [C101Component, SanitizeHtmlPipe],
  exports: [SanitizeHtmlPipe],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class C101Module {
}
