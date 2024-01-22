import {afterNextRender, Component} from '@angular/core';
import {Meta, Title} from '@angular/platform-browser';

import {B101Service} from "./b101.service";
import {PkgContentClass} from '../../../model/pkgContentClass';

@Component({
  selector: 'app-b101',
  templateUrl: './b101.component.html',
})
export class B101Component {

  proId: string = 'b101';
  content: PkgContentClass[] = [];

  constructor(
    public service: B101Service,
    private meta: Meta,
    public title: Title
  ) {
    const ogTitle: string = 'jSql | JGo';
    const description: string = 'JSql designed on the basis of mysql, go-mssqldb, godror and pq.';
    const keywords: string = 'jsql,mssql,mysql,oracle';
    const ogUrl: string = 'https://jgo.dev/pkg/jsql';
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
    });
  }
}
