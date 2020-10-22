# pure python gene sequencing


def gene_slicer(genetic_code):
    if set(genetic_code.upper()).issubset(['A', 'C', 'G', 'T']):
        genes = {}
        while len(genetic_code) >= 3:
            genes.update({genetic_code[:3]: (genes.get(genetic_code[:3]) or 0) + 1})
            genetic_code = genetic_code[1:]
        print(genes)
    else:
        print("Please input a valid genetic sequence")


gene_slicer("AATCCGCTAG")
