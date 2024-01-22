import {NgModule} from '@angular/core';
import {BrowserModule, provideClientHydration} from '@angular/platform-browser';
import {HttpClientModule} from "@angular/common/http";

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';

import {LoadingComponent} from './views/loading/loading.component';
import {MenuComponent} from './views/menu/menu.component';
import {FooterComponent} from './views/footer/footer.component';
import {SharedComponent} from './views/shared/shared.component';
import {ErrorComponent} from './views/error/error.component';

import {AModule} from './views/A/a.module';
import {BModule} from './views/B/b.module';
import {CModule} from './views/C/c.module';

@NgModule({
  declarations: [
    AppComponent,
    LoadingComponent,
    MenuComponent,
    FooterComponent,
    SharedComponent,
    ErrorComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    AModule,
    BModule,
    CModule,
  ],
  providers: [
    provideClientHydration()
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
