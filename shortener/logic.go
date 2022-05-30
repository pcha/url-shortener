package shortener

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrRedirecNotFound = errors.New("Redirect Not Found")
	ErrRedirecInvalid  = errors.New("RedirectInvalid")
)

type redirectService struct {
	redirectRepository RedirectRepository
}

func NewRedirectService(repository RedirectRepository) RedirectService {
	return &redirectService{
		repository,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepository.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirecInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepository.Store(redirect)
}
