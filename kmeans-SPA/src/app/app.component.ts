import { Component, OnInit } from '@angular/core';
import { ClusteringService } from './_services/clustering.service';
import { Diagnostic } from './_models/diagnostic';
import { GroupedData } from './_models/grouped-data';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{

  diagnostics: Diagnostic[];
  groupedData?: GroupedData;


  constructor(private clusteringService: ClusteringService) {}

  ngOnInit(): void {

  }

  agruparDatos(): void{
    this.clusteringService.getClusters().subscribe( (response) => {
      this.groupedData = response;
    }, error => {
      console.log(error);
    });
  }


}
