import { Component } from '@angular/core';
import { ApiService } from './api/services/api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
    title = 'ui';

    frequency: string

    constructor(private apiService: ApiService) {}

    ngOnInit(): void {
        this.getFrequency();
    }

    powerOn(): void {
        this.apiService.powerstatPut({"body": {"status": 1}}).subscribe(resp => console.log(resp));
        this.getFrequency();
    }

    powerOff(): void {
        this.apiService.powerstatPut({"body": {"status": 0}}).subscribe(resp => console.log(resp));
    }

    getFrequency(): void {
        this.apiService.frequencyGet().subscribe(resp => (
            this.frequency = resp.data.frequency
        ))
    }

    setFrequency(f: string): void {
        console.log("Updating freq "+f)
        this.apiService.frequencyPut({"body": {"frequency": f}}).subscribe(function(resp){
            console.log(resp);
            this.frequency = '123456789'
        });
    }
}
