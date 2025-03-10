// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztables

import (
	"fmt"
	"strings"
)

// NewServiceClientFromConnectionString creates a new ServiceClient struct from a connection string. The connection
// string must contain either an account name and account key or an account name and a shared access signature.
func NewServiceClientFromConnectionString(connectionString string, options *ClientOptions) (*ServiceClient, error) {
	endpoint, credential, err := parseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	if credential == nil {
		return NewServiceClientWithNoCredential(endpoint, options)
	}
	return NewServiceClientWithSharedKey(endpoint, credential, options)
}

// convertConnStrToMap converts a connection string (in format key1=value1;key2=value2;key3=value3;) into a map of key-value pairs
func convertConnStrToMap(connStr string) (map[string]string, error) {
	ret := make(map[string]string)
	connStr = strings.TrimRight(connStr, ";")

	splitString := strings.Split(connStr, ";")
	if len(splitString) == 0 {
		return ret, errConnectionString
	}
	for _, stringPart := range splitString {
		parts := strings.SplitN(stringPart, "=", 2)
		if len(parts) != 2 {
			return ret, errConnectionString
		}
		ret[parts[0]] = parts[1]
	}
	return ret, nil
}

// parseConnectionString parses a connection string into a service URL and a SharedKeyCredential or a service url with the
// SharedAccessSignature combined.
func parseConnectionString(connStr string) (string, *SharedKeyCredential, error) {
	var serviceURL string
	var cred *SharedKeyCredential

	defaultScheme := "https"
	defaultSuffix := "core.windows.net"

	connStrMap, err := convertConnStrToMap(connStr)
	if err != nil {
		return "", nil, err
	}

	accountName, ok := connStrMap["AccountName"]
	if !ok {
		return "", nil, errConnectionString
	}
	accountKey, ok := connStrMap["AccountKey"]
	if !ok {
		sharedAccessSignature, ok := connStrMap["SharedAccessSignature"]
		if !ok {
			return "", nil, errConnectionString
		}
		return fmt.Sprintf("%v://%v.table.%v/?%v", defaultScheme, accountName, defaultSuffix, sharedAccessSignature), nil, nil
	}

	protocol, ok := connStrMap["DefaultEndpointsProtocol"]
	if !ok {
		protocol = defaultScheme
	}

	suffix, ok := connStrMap["EndpointSuffix"]
	if !ok {
		suffix = defaultSuffix
	}

	tableEndpoint, ok := connStrMap["TableEndpoint"]
	if ok {
		cred, err = NewSharedKeyCredential(accountName, accountKey)
		return tableEndpoint, cred, err
	}
	serviceURL = fmt.Sprintf("%v://%v.table.%v", protocol, accountName, suffix)

	cred, err = NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return "", nil, err
	}

	return serviceURL, cred, nil
}
