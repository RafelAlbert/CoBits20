package main

import "time"

type AlumneAula struct {
  Aula string
  Nom  string
  X    float64
  Y    float64
}

type Alumne struct {
  Nom string
  CoordX  float64
  CoordY  float64
}

type GAlumne map[string]map[string]time.Time

type Aules map[string][]Alumne

var aules Aules = make(Aules)

var estatAlumne map[string]string = make(map[string]string)

var currentAlumne AlumneAula

var contactes GAlumne = make(map[string]map[string]time.Time)

