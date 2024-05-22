package users

type Hombre struct {
	Persona
}

func (h *Hombre) Respirar() { h.respirando = true }
func (h *Hombre) Pensar() { h.pensando = true }
func (h *Hombre) Comer() { h.comiendo = true }
func (h *Hombre) Sexo() string { return "hombre" }
