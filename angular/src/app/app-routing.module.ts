import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

import {SharedComponent} from './views/shared/shared.component';
import {ErrorComponent} from './views/error/error.component'

import {A101Component} from './views/A/a101/a101.component';
import {B101Component} from './views/B/b101/b101.component';
import {B102Component} from './views/B/b102/b102.component';
import {B103Component} from './views/B/b103/b103.component';
import {C101Component} from './views/C/c101/c101.component';

const routes: Routes = [
  {path: '', redirectTo: 'home', pathMatch: 'full'},
  {
    path: '', component: SharedComponent,
    children: [
      {path: 'home', component: A101Component},
      {path: 'pkg/jsql', component: B101Component},
      {path: 'pkg/jcron', component: B102Component},
      {path: 'pkg/jlog', component: B103Component},
      {path: 'update', component: C101Component},
    ]
  },
  {path: '**', component: ErrorComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
