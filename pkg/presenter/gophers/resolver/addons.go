package graphResolver

import (
	domainAccess "github.com/travelgateX/presenters-benchmark/pkg/access"
)

const (
	XoParamPrefix  = "xo_"
	XodParamPrefix = "xod_"
)

type AddOnsResolver struct {
	addons *AddonMap
}

type AddonMap struct {
	addons map[string]map[string]string
}

func (a *AddonMap) AddParam(addonType string, p domainAccess.Parameter) {
	if a.addons == nil {
		a.addons = make(map[string]map[string]string)
	}
	if _, ok := a.addons[addonType]; !ok {
		a.addons[addonType] = make(map[string]string)
	}
	a.addons[addonType][p.Key[len(addonType):]] = p.Value
}

func (r *AddOnsResolver) Distribute() *Json {
	ret := ""
	if r.addons != nil && r.addons.addons[XoParamPrefix] != nil && r.addons.addons[XoParamPrefix]["Breakdown"] != "" {
		ret = r.addons.addons[XoParamPrefix]["Breakdown"]
	}
	if ret == "" && r.addons != nil && r.addons.addons[XodParamPrefix] != nil && r.addons.addons[XodParamPrefix]["Breakdown"] != "" {
		ret = r.addons.addons[XodParamPrefix]["Breakdown"]
	}
	if ret == "" {
		return nil
	}
	tmp := Json(ret)
	return &tmp
}

func (r *AddOnsResolver) Distribution() *[]*AddOnResolver {
	if r.addons != nil {
		if dic, ok := r.addons.addons[XodParamPrefix]; ok {
			addons := make([]*AddOnResolver, 0, len(dic))
			for key, value := range dic {
				addons = append(addons, &AddOnResolver{key: key, value: value})
			}
			return &addons
		}
	}
	return nil
}
