import os


def load_code_changes(filename):
    code_changes = {}
    with open(filename, encoding='utf-8') as f:
        for line in f:
            cells = line.strip().split(',')
            old = cells[1]
            new = cells[2]
            if old not in code_changes:
                code_changes[old] = {}
            code_changes[old]['new'] = f'{new}'
            dists = cells[3].split(' -> ')
            code_changes[old]['old_district'] = f'{dists[0]}'
            code_changes[old]['new_district'] = f'{dists[0]}'
            if len(dists) > 1:
                code_changes[old]['new_district'] = f'{dists[1]}'
    return code_changes


def load_name_changes(filename):
    name_changes = {}
    with open(filename, encoding='utf-8') as f:
        for line in f:
            cells = line.strip().split(',')
            change_year = int(cells[0].split('-')[1])
            code = cells[1]
            if change_year not in name_changes:
                name_changes[change_year] = {}
            name_changes[change_year][code] = code
    return name_changes


def load_merges(filename):
    merges = {}
    with open(filename, encoding='utf-8') as f:
        for line in f:
            [_, old, new] = line.strip().split(',')
            merges[old] = new
    return merges


def load_splits(filename):
    splits = {}
    with open(filename, encoding='utf-8') as f:
        for line in f:
            cells = line.strip().split(',')
            change_year = int(cells[0])
            old = cells[1]
            if change_year not in splits:
                splits[change_year] = {}
            splits[change_year][old] = cells[2:]
    return splits


def path_wrapper(filename):
    return os.path.join(os.path.dirname(__file__), filename)


merges = load_merges(path_wrapper('csv/code-merges.csv'))
splits = load_splits(path_wrapper('csv/code-splits.csv'))
code_changes = load_code_changes(path_wrapper('csv/code-changes.csv'))
name_changes = load_name_changes(path_wrapper('csv/name-changes.csv'))


def sketch_code_changes():
    js_str = '{'
    for k, v in code_changes.items():
        js_str += f'\n\t"{k}":' + '{'
        js_str += f'\n\t\t"new_code":"{v["new"]}",'
        js_str += f'\n\t\t"new_district":"{v["new_district"]}",'
        js_str += f'\n\t\t"old_district":"{v["old_district"]}"'
        js_str += '\n\t},'
    js_str = js_str[:-1] + '\n}'
    return js_str


def sketch_code_merges():
    js_str = '{'
    for k, v in merges.items():
        js_str += f'\n\t"{k}":"{v}",'
    js_str = js_str[:-1] + '\n}'
    return js_str


if __name__ == "__main__":
    with open('json/address_code_changes.json', "w+", encoding='utf-8') as f:
        f.write(sketch_code_changes())
    with open('json/address_code_merges.json', "w+", encoding='utf-8') as f:
        f.write(sketch_code_merges())
