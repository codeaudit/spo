import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class PurchaseService {

  private purchaseOrders: Subject<any[]> = new BehaviorSubject<any[]>([]);
  // private purchaseUrl: string = 'https://event.spaco.io/api/';
  private purchaseUrl: string = 'http://121.42.24.199:7071/api/';
  // private purchaseUrl: string = '/teller/';

  private purchaseTokenTypes: Subject<TokenModel[]> = new BehaviorSubject<TokenModel[]>([]);
  
  //buy types
  //private purchaseTokenTypes :Subject<any[]> = new BehaviorSubject<any[]>([]);

  constructor(
    private http: Http,
  ) {
    this.retrievePurchaseOrders();
    this.retrievePurchaseTokenTypes();
  }

  all() {
    return this.purchaseOrders.asObservable();
  }

  generateTokenAddress(address: string,coinType) {
    return this.post('bind', { address: address,plan_coin_type:coinType })
      .do(response => {
        this.purchaseOrders.first().subscribe(orders => {
          let index = orders.findIndex(order => order.address === address);
          if (index === -1) {
            orders.push({address: address, addresses: []});
            index = orders.length - 1;
          }
          const timestamp = Math.floor(Date.now() / 1000);
          orders[index].addresses.unshift({
            btc: response.btc_address,
            status: 'waiting_deposit',
            created: timestamp,
            updated: timestamp,
          });
          this.updatePurchaseOrders(orders)
        });
      });
  }

  scanTokenAddress(address: string) {
    return this.get('status?skyaddr=' + address).do(response => {
      this.purchaseOrders.first().subscribe(orders => {
        let index = orders.findIndex(order => order.address === address);
        // Sort addresses ascending by creation date to match teller status response
        orders[index].addresses.sort((a, b) =>  b.created - a.created);
        for (const btcAddress of orders[index].addresses) {
          // Splice last status to assign this to the latest known order
          const status = response.statuses.splice(-1,1)[0];
          btcAddress.status = status.status;
          btcAddress.updated = status.update_at;
        }

        this.updatePurchaseOrders(orders)
      });
    });
  }


  generate(address: string) {
    return this.post('bind', { skyaddr: address })
      .do(response => {
        this.purchaseOrders.first().subscribe(orders => {
          let index = orders.findIndex(order => order.address === address);
          if (index === -1) {
            orders.push({address: address, addresses: []});
            index = orders.length - 1;
          }
          const timestamp = Math.floor(Date.now() / 1000);
          orders[index].addresses.unshift({
            btc: response.btc_address,
            status: 'waiting_deposit',
            created: timestamp,
            updated: timestamp,
          });
          this.updatePurchaseOrders(orders)
        });
      });
  }

  scan(address: string) {
    return this.get('status?skyaddr=' + address).do(response => {
      this.purchaseOrders.first().subscribe(orders => {
        let index = orders.findIndex(order => order.address === address);
        // Sort addresses ascending by creation date to match teller status response
        orders[index].addresses.sort((a, b) =>  b.created - a.created);
        for (const btcAddress of orders[index].addresses) {
          // Splice last status to assign this to the latest known order
          const status = response.statuses.splice(-1,1)[0];
          btcAddress.status = status.status;
          btcAddress.updated = status.update_at;
        }

        this.updatePurchaseOrders(orders)
      });
    });
  }
  
  alltokens(): Observable<TokenModel[]> {
    return this.purchaseTokenTypes.asObservable();
  }

  getSupportedTokens() {
    return this.get('tokens').do(response=>{
        this.purchaseTokenTypes.first().subscribe(tokens=> {
            this.updateTokenTypes(tokens)
        });
    });    
  }
    

  private get(url) {
    return this.http.get(this.purchaseUrl + url)
      .map((res: any) => res.json())
  }

  private post(url, parameters = {}) {
    return this.http.post(this.purchaseUrl + url, parameters)
      .map((res: any) => res.json())
  }

  private retrievePurchaseOrders() {
    const orders = JSON.parse(window.localStorage.getItem('purchaseOrders'));
    if (orders) {
      this.purchaseOrders.next(orders);
    }
  }

  private retrievePurchaseTokenTypes() {
    const tokenTypes = JSON.parse(window.localStorage.getItem('purchaseTokenTypes'));
    if(tokenTypes) {
      this.purchaseTokenTypes.next(tokenTypes);
    }

  }

  private updatePurchaseOrders(collection: any[]) {
    this.purchaseOrders.next(collection);
    window.localStorage.setItem('purchaseOrders', JSON.stringify(collection));
  }

  private updateTokenTypes(collection: any[]) {
    this.purchaseTokenTypes.next(collection);
    window.localStorage.setItem('purchaseTokenTypes', JSON.stringify(collection));
  }
}
