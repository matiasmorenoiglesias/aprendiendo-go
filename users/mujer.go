package users

type Mujer struct {
	Persona
	EsMadre bool
}

func (h *Mujer) Respirar() { h.respirando = true }
func (h *Mujer) Pensar() { h.pensando = true }
func (h *Mujer) Comer() { h.comiendo = true }
func (h *Mujer) Sexo() string { return "mujer" }
