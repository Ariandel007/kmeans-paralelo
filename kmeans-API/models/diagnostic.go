package models

type Diagnostic struct {
	Age uint `json:"age"` //age
	Sex uint `json:"sex"` //sex
	Cp uint `json:"cp"` //chest pain(dolor en el pecho
	Trestbps uint `json:"trestbps"`// presión arterial en reposo
	Chol uint `json:"chol"` // colesterol serico en mg/dl
	Fbs uint `json:"fbs"` //// azúcar en sangre en ayunas> 120 mg / dl
	Restecg uint `json:"restecg"` // resultados electrocardiográficos en reposo (valores 0,1,2)
	Thalach uint `json:"thalach"` // frecuencia cardíaca máxima alcanzada
	Exang uint `json:"exang"` //exercise induced angina
	Oldpeak float64 `json:"oldpeak"` // oldpeak = depresión inducida por el ejercicio relativo al descanso
	Slope uint `json:"slope"` // la pendiente del segmento de ejercicio pico
	Ca uint `json:"ca"` // número de vasos principales (0-3) coloreados por flourosopía
	Thal uint `json:"thal"` // thal: 3 = normal; 6 = defecto fijo; 7 = defecto reversible
	Target uint `json:"target"` //esta enfermo o no, 0 o 1

}