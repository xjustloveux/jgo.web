import {Component, Input} from '@angular/core';

import {PkgContentClass} from '../../model/pkgContentClass';

@Component({
  selector: 'app-pkg-content',
  templateUrl: './pkg-content.component.html',
})
export class PkgContentComponent {

  @Input('content') content: PkgContentClass[] = [];
}
