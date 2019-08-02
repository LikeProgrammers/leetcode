package groupAnagrams

import (
    "sort"
    "strings"
)

func groupAnagrams(strs []string) [][]string {
    m_group := map[string][]string{}
    group := [][]string{}
    for _, v := range strs {
        s := strings.Split(v, "")
        sort.Strings(s)
        ss := strings.Join(s, "")
        m_group[ss] = append(m_group[ss], v)
    }
    for _, v := range m_group {
        group = append(group, v)
    }
    return group
}
