import { getComponentStateName } from '../../../store/states';

export default function() {
  this.componentId = getComponentStateName(this._riot_id);
  this.cardType = null;
  // 縦横何個分のcellを使うのか。
  this.columnSize = 'columnSpreadSmall';
  this.rowSize = 'rowSpreadSmall';
  switch (this.opts.def.style) {
  case 'chart':
    // v1ではHighchartが使用されていたがライセンス上の問題で使用しないことに。
    // `chart` = highchart使用なのでサポート外である旨を示す。
    this.cardType = 'components-page-unsupported';
    this.columnSize = 'columnSpreadSmall';
    this.rowSize = 'rowSpreadMedium';
    break;
  case 'graph-scatterplot':
  case 'graph-line':
  case 'graph-bar':
  case 'graph-horizontal-bar':
  case 'graph-stacked-bar':
  case 'graph-horizontal-stacked-bar':
  case 'graph-stacked-area':
    this.cardType = 'components-page-chart';
    this.columnSize = 'columnSpreadSmall';
    this.rowSize = 'rowSpreadMedium';
    break;
  case 'table':
    this.cardType = 'components-page-table';
    this.columnSize = 'columnSpreadFull';
    this.rowSize = 'rowSpreadLarge';
    break;
  case 'number':
    this.cardType = 'components-page-number';
    this.columnSize = 'columnSpreadSmall';
    this.rowSize = 'rowSpreadSmall';
    break;
  case 'explorer':
    // explorerはwyswyg等から使用されるのでsrc/components配下にソースを置いている。
    this.cardType = 'explorer';
    this.columnSize = 'columnSpreadFull';
    this.rowSize = 'rowSpreadLarge';
    break;
  default:
    break;
  }
}
