import { Component, OnInit } from '@angular/core';
import { ClusteringService } from './_services/clustering.service';
import { Diagnostic } from './_models/diagnostic';
import { GroupedData } from './_models/grouped-data';
import { ListDiagnostic } from './_models/list-diagnostic';
import { DiagnosticService } from './_services/diagnostic.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{

  diagnostics: ListDiagnostic;
  groupedData?: GroupedData;
  k: number;
  seccion = 1;


  constructor(private clusteringService: ClusteringService, private diagnosticService: DiagnosticService) {}

  ngOnInit(): void {
    this.obtenerListaDiagnosticos();
  }

  agruparDatos(): void{
    console.log(this.k);
    if ( this.k <= 0 ) {
      return;
    }

    this.clusteringService.getClusters(this.k).subscribe( (response) => {
      this.groupedData = response;
    }, error => {
      console.log(error);
    });
  }

  obtenerListaDiagnosticos(): void {
    this.diagnosticService.getDiagnostics().subscribe( (response) => {
      this.diagnostics = response;
    }, error => {
      console.log(error);
    });
  }

  setSeccion(n: number) {
    this.seccion = n;
  }

}
