import re


def gene_slicer(genetic_code):
    if re.fullmatch('[ACTG]*', genetic_code) is not None:
        genes = []
        for i in range(len(genetic_code) - 2):
            current_gene = genetic_code[i:i + 3]
            if current_gene not in genes:
                genes.append(current_gene)
        for gene in genes:
            print(gene + ":" + str(genetic_code.count(gene)))
    else:
        print("please input a valid genetic sequence")


gene_slicer("AATCCGCTAG")
