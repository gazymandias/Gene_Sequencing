# pure python gene sequencing


def find_start_codon(genetic_code, start_codon='ATG'):
    if start_codon not in genetic_code: 
        print("No valid start codon")
        return
    else:
       return genetic_code.find(start_codon)

    
def find_stop_codon(genetic_code, start_codon, stop_codons=['TAA', 'TGA', 'TAG']):
    if any(x in genetic_code for x in stop_codons):
        stop_codon_positions = [genetic_code.find(x) for x in stop_codons if 0 < genetic_code.find(x) > start_codon]
        if stop_codon_positions:
            return min(stop_codon_positions)
        else:
            print("No valid stop codons AFTER start codon")
            return
    else:
        print("No valid stop codons")
        return
        

def check_valid_code(genetic_code):
    if not set(genetic_code.upper()).issubset(['A', 'C', 'G', 'T']):
        print("Please input a valid genetic sequence")
        return False
    else:
        return True


def sequence_codons(genetic_code, start_codon_position, stop_codon_position):
    genetic_code = genetic_code[start_codon_position:stop_codon_position]
    print(f"sequencing the following subset of code with start codon position ({start_codon_position}) and stop codon position ({stop_codon_position}): {genetic_code}")
    genes = {}
    while len(genetic_code) >= 3:
        genes.update({genetic_code[:3]: (genes.get(genetic_code[:3]) or 0) + 1})
        genetic_code = genetic_code[3:]
    print(genes)
    return genes


def main():
    try:
        genetic_code = "cgcagacgcgcaagcagattttataaagccaatgtatatacaactcctaccaggccaaacaaggtccgcgcacacgaaacctggaaaactttgttgagtcggc".upper()
        is_valid = check_valid_code(genetic_code=genetic_code)
        if is_valid:
            start_codon_position = find_start_codon(genetic_code)
            stop_codon_position = find_stop_codon(genetic_code, start_codon_position)
            if start_codon_position and stop_codon_position:
                sequence_codons(genetic_code, start_codon_position, stop_codon_position)
    except Exception as SuperGenericException:
        print(SuperGenericException)


if __name__ == '__main__':
    main()
