import {Component} from '@angular/core';
import {Router} from '@angular/router';
import {Meta, Title} from '@angular/platform-browser';

@Component({
  selector: 'app-error',
  templateUrl: './error.component.html',
})
export class ErrorComponent {

  constructor(
    private meta: Meta,
    public title: Title,
    private router: Router
  ) {
    const ogTitle: string = 'Error | JGo';
    const description: string = 'JGo provides an easier configuration for writing sql, log, and cron jobs.';
    const keywords: string = 'jgo,golang,cron,sql,log,job,schedule,mssql,mysql,oracle';
    const ogUrl: string = 'https://jgo.dev/pkg/error';
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
    this.title.setTitle('Error | JGo');
  }

  bnClick(): void {
    this.router.navigate(['/home']).then(() => {
    });
  }
}
