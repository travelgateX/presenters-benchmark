package graphResolver

import (
	domainAccess "hub-aggregator/common/domain/access"
	"hub-aggregator/common/graphql"
	"hub-aggregator/common/hotel-common/token/optionId"
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

func (r *AddOnsResolver) Distribute() *graphql.Json {
	ret := ""
	if r.addons != nil && r.addons.addons[optionId.XoParamPrefix] != nil && r.addons.addons[optionId.XoParamPrefix]["Breakdown"] != "" {
		ret = r.addons.addons[optionId.XoParamPrefix]["Breakdown"]
	}
	if ret == "" && r.addons != nil && r.addons.addons[optionId.XodParamPrefix] != nil && r.addons.addons[optionId.XodParamPrefix]["Breakdown"] != "" {
		ret = r.addons.addons[optionId.XodParamPrefix]["Breakdown"]
	}
	if ret == "" {
		return nil
	}
	tmp := graphql.Json(ret)
	return &tmp
}

func (r *AddOnsResolver) Distribution() *[]*AddOnResolver {
	if r.addons != nil {
		if dic, ok := r.addons.addons[optionId.XodParamPrefix]; ok {
			addons := make([]*AddOnResolver, 0, len(dic))
			for key, value := range dic {
				addons = append(addons, &AddOnResolver{key: key, value: value})
			}
			return &addons
		}
	}
	return nil
}
