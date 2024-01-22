import {NgModule, Pipe, PipeTransform} from '@angular/core';
import {CommonModule} from '@angular/common';
import {DomSanitizer, SafeHtml} from '@angular/platform-browser';
import {PkgContentComponent} from './pkg-content.component';

@Pipe({name: 'sanitizeHtml'})
export class SanitizeHtmlPipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) {
  }

  transform(v: string): SafeHtml {
    return this.sanitizer.bypassSecurityTrustHtml(v);
  }
}

@NgModule({
  declarations: [PkgContentComponent, SanitizeHtmlPipe],
  exports: [
    PkgContentComponent,
    SanitizeHtmlPipe
  ],
  imports: [CommonModule]
})
export class PkgContentModule {
}
