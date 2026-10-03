package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.StatusMove;
import ru.ifmo.se.pokemon.Type;

final public class PoisonPowder extends StatusMove {
  public PoisonPowder() {
    super(Type.POISON, 0, 0.75);
  }

  @Override
  public void applyOppEffects(Pokemon p) {
    Effect.poison(p);
  }

  @Override
  public String describe() {
    return "uses ability \"Poison Powder\"";
  }
}
