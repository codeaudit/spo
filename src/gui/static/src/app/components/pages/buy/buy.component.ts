import { Component } from '@angular/core';
import { PurchaseService } from '../../../services/purchase.service';
import { MdDialog, MdDialogConfig, MdDialogRef } from '@angular/material';
import { AddDepositAddressComponent } from './add-deposit-address/add-deposit-address.component';
import { config } from '../../../app.config';

@Component({
  selector: 'app-buy',
  templateUrl: './buy.component.html',
  styleUrls: ['./buy.component.css']
})
export class BuyComponent {

  orders = [];
  otcEnabled: boolean;
  scanning = false;
  supportedTokens = [];

  constructor(
    public purchaseService: PurchaseService,
    private dialog: MdDialog,
  ) {
    this.otcEnabled = config.otcEnabled;
    this.load();
  }

  addDepositAddress(token) {
    //console.log(token)
    const config = new MdDialogConfig();
    config.width = '500px';
    let dialogRef:MdDialogRef<AddDepositAddressComponent> = this.dialog.open(AddDepositAddressComponent, config);
    dialogRef.componentInstance.tokenType = token;
  }

  searchDepositAddress(address: string) {
    this.scanning = true;
    this.purchaseService.scan(address).subscribe(() => {
      this.disableScanning();
    }, error => {
      this.disableScanning();
    });
  }

  load() {
    //this.purchaseService.getSupportedTokens();
    this.getSupportedTokens();
  }
  
  getSupportedTokens(){
    this.purchaseService.getSupportedTokens().subscribe(result => {
      //console.log(result);
      if (result.code != 0) {
        alert(result.errmsg);
        return;
      }
      this.supportedTokens = result.data;        
    })    
  }
  
  private disableScanning()
  {
    setTimeout(() => this.scanning = false, 1000);
  }
  ngOnInit(): void {
    
  }
}
