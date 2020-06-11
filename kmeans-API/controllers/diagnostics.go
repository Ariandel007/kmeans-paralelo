package controllers


//importamos gin que es un mini-framework que nos da lo necesario para realizar una API REST
import (
	"../kmeans"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RealizarClustering(c *gin.Context) {

	data := []kmeans.Punto{}


	var diagnostics []models.Diagnostic
	models.DB.Find(&diagnostics)

	for _, diagnostic := range diagnostics{
		data = append(data, kmeans.Punto{[]float64{
			float64(diagnostic.Age),
			float64(diagnostic.Sex),
			float64(diagnostic.Cp),
			float64(diagnostic.Trestbps),
			float64(diagnostic.Chol),
			float64(diagnostic.Fbs),
			float64(diagnostic.Restecg),
			float64(diagnostic.Thalach),
			float64(diagnostic.Exang),
			diagnostic.Oldpeak,
			float64(diagnostic.Slope),
			float64(diagnostic.Ca),
			float64(diagnostic.Thal),
			float64(diagnostic.Target)}})
	}

	var k uint64 = 5
	var clusters = kmeans.KMEANS(data, k, 1)

	var listRes [][]models.Diagnostic

	if len(clusters) > 0 {
		for _, c := range clusters{
			var resultado []models.Diagnostic
			for _, p := range c.Puntos{
				var d models.Diagnostic
				d.Age = uint(p.Entrada[0])
				d.Sex = uint(p.Entrada[1])
				d.Cp = uint(p.Entrada[2])
				d.Trestbps = uint(p.Entrada[3])
				d.Chol = uint(p.Entrada[4])
				d.Fbs = uint(p.Entrada[5])
				d.Restecg = uint(p.Entrada[6])
				d.Thalach = uint(p.Entrada[7])
				d.Exang = uint(p.Entrada[8])
				d.Oldpeak = p.Entrada[9]
				d.Slope = uint(p.Entrada[10])
				d.Ca = uint(p.Entrada[11])
				d.Thal = uint(p.Entrada[12])
				d.Target = uint(p.Entrada[13])

				resultado = append(resultado, d)
			}
			listRes = append(listRes, resultado)
		}
	}


	c.JSON(http.StatusOK, gin.H{"data": listRes})

}