package parser

import (
    "github.com/hashicorp/hcl/v2/hclsimple"
    "hcl-redis/types"
)

func Parse(route string) (*types.Document, error) {
    var doc types.Document
    err := hclsimple.DecodeFile(route, nil, &doc)
    if err != nil {
        return nil, err
    }
    return &doc, nil
}

func ParseString(content string) (*types.Document, error) {
    var doc types.Document
    err := hclsimple.Decode("input.hcl", []byte(content), nil, &doc)
    if err != nil {
        return nil, err
    }
    return &doc, nil
}