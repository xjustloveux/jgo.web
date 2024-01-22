import {afterNextRender, Component} from '@angular/core';
import {Meta, Title} from '@angular/platform-browser';

import {B103Service} from './b103.service';
import {PkgContentClass} from '../../../model/pkgContentClass';
import {LogFileClass} from '../../../model/logFileClass';

@Component({
  selector: 'app-b103',
  templateUrl: './b103.component.html',
})
export class B103Component {

  proId: string = 'b103';
  content: PkgContentClass[] = [];
  log: LogFileClass | undefined;

  constructor(
    public service: B103Service,
    private meta: Meta,
    public title: Title
  ) {
    const ogTitle: string = 'jLog | JGo';
    const description: string = 'jLog is a log library written in Go.';
    const keywords: string = 'jlog,log,library';
    const ogUrl: string = 'https://jgo.dev/pkg/jlog';
    const metaList: any = [
      {name: 'description', content: description},
      {name: 'keywords', content: keywords},
      {name: 'twitter:title', content: ogTitle},
      {name: 'twitter:description', content: description},
      {property: 'og:title', content: ogTitle},
      {property: 'og:description', content: description},
      {property: 'og:url', content: ogUrl},
    ];
    for (const item of metaList) {
      this.meta.updateTag(item);
    }
    this.title.setTitle(ogTitle);
    afterNextRender(() => {
      this.service.queryContent(this.proId).then((res) => {
        this.content = res;
      });
      this.service.queryLog().then((res) => {
        this.log = res;
      });
    });
  }
}
